/*
 * what: diconnecting
 * why:  we want to disconnect a signal from a slot
 * how:
 *      use the diconnect funciton
 *      this alos talk about conncetion types
 *
 */

#include <QCoreApplication>
#include <QDebug>
#include <QTextStream>
#include <iostream>

#include "radio.h"
#include "station.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    Radio boombox;
    //Station* channels[3];

    //todo: replace with a collection

    std::vector<Station*> channels;
    channels.push_back( new Station(&boombox,  94, "Rock & Roll") );
    channels.push_back( new Station(&boombox,  87, "Hip Hop") );
    channels.push_back( new Station(&boombox, 104, "News") );

    boombox.connect( &boombox, &Radio::quit, &a, &QCoreApplication::quit );

    for( auto &pc : channels )
    {
        boombox.connect(pc, &Station::send, &boombox, &Radio::listen);
    }

    // for(int i = 0; i<3; i++)
    // {
    //     Station* channel = channels[i];
    //     boombox.connect(channel, &Station::send, &boombox, &Radio::listen );
    // }

    do
    {
        qInfo() << "Enter on, off, test, or quit";
        QTextStream qtin(stdin);
        QString line = qtin.readLine().trimmed().toUpper();

        if( line == "T" || line.mid(0,3)=="TEST")
        {
            qInfo() << "Testing";
            for( auto &pc : channels )
            {
                pc->broadcast("Broadcasting live!");
            }
        }
        else if(line == "1" || line.mid(0,1)=="ON")
        {
            qInfo() << "Turning the radio on";
            for( auto &pc : channels )
            {
                boombox.connect(pc, &Station::send, &boombox, &Radio::listen);
            }
        }
        else if(line == "0" || line == "OFF" || line.mid(0,2)=="OFF")
        {
            qInfo() << "Turning the radio off";
            for( auto &pc : channels )
            {
                boombox.disconnect(pc, &Station::send, &boombox, &Radio::listen);
            }
        }
        else if(line == "Q" || line == "QUIT" || line.mid(0,3)=="QUIT")
        {
            qInfo() << "Quitting";
            emit boombox.quit();
            break;
        }
        else
        {
            qInfo() << "Say WHAT?";
            qInfo() << "Don't understand command: " << line;
        }


    } while(true);

    return a.exec();
}
