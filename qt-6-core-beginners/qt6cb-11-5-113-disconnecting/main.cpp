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

    struct pallet{
        Station * _pc;
        bool _isConnected;
        pallet(Station* pc):_pc(pc),_isConnected(false){}
        QString toString(){return _pc!=nullptr ? _pc->toString() : "empty";}
        void broadcast(const QString& msg){ if(_pc!=nullptr) _pc->broadcast(msg);}
    };
    std::vector<pallet> channels;
    channels.push_back( pallet{new Station(&boombox,  94, "Rock & Roll")} );
    channels.push_back( pallet{new Station(&boombox,  87, "Hip Hop")} );
    channels.push_back( pallet{new Station(&boombox, 104, "News") });

    boombox.connect( &boombox, &Radio::quit, &a, &QCoreApplication::quit, Qt::QueuedConnection );

    for( auto &pallet : channels )
    {
        boombox.connect(pallet._pc, &Station::send, &boombox, &Radio::listen);
    }

    // for(int i = 0; i<3; i++)
    // {
    //     Station* channel = channels[i];
    //     boombox.connect(channel, &Station::send, &boombox, &Radio::listen );
    // }

    do
    {
        qInfo() << "Enter on, off, test, list, or quit";
        QTextStream qtin(stdin);
        QString line = qtin.readLine().trimmed().toUpper();
        const QString
            sHELP{"HELP"},
            sLIST{"LIST"},
            sTEST{"TEST"},
            sON{"ON"},
            sCONNECT("CONNECT"),
            sOFF{"OFF"},
            sDISCONNECT("DISCONNECT"),
            sQUIT{"QUIT"};

        if( sTEST.contains(line) || line == "3")
        {
            qInfo() << "Testing";
            for( auto &pallet : channels )
            {
                pallet._pc->broadcast("Broadcasting live!");
            }
            qInfo() << "";
        }
        else if(sLIST.contains(line) || line == "2" )
        {
            qInfo() << "List of channels:";
            for( auto &pallet : channels )
            {
                qInfo() <<  pallet._pc->channel << " " << pallet._pc->name
                        << ", isConnected?: " << pallet._isConnected;
            }
            qInfo() << "";

        }
        else if(sON.contains(line) || sCONNECT.contains(line) || line == "1" )
        {
            qInfo();
            qInfo() << "Turning the radio on";
            for( auto &pallet : channels )
            {
                if(pallet._isConnected){
                    qInfo() << "  channel: " << pallet._pc->toString() <<" is already connected.";
                }else{
                    boombox.connect(pallet._pc, &Station::send, &boombox, &Radio::listen);
                    qInfo() << "  connecting: " << pallet._pc->toString();
                    pallet._isConnected = true;
                }
                boombox.connect(pallet._pc, &Station::send, &boombox, &Radio::listen);
            }
            qInfo() << "";
        }
        else if(sOFF.contains(line) || sDISCONNECT.contains(line)|| line == "0") // off
        {
            qInfo() << "Turning the radio off";
            for( auto &pallet : channels )
            {
                qInfo() << "  disconnecting: " << pallet._pc->toString();

                boombox.disconnect(pallet._pc, &Station::send, &boombox, &Radio::listen);
                pallet._isConnected = false;

            }
            qInfo() << "";
        }
        else if(sQUIT.contains(line) || line == "4")
        {
            qInfo() << "Quitting";
            emit boombox.quit();
            break;
        }
        else if(sHELP.contains(line) || line == "9")
        {
            qInfo() << "";
            qInfo() << "commands:";
            qInfo() << "  help: displays this text";
            qInfo() << "    on:   connects";
            qInfo() << "   off: diconnects";
            qInfo() << "  test: displays test msg";
            qInfo() << "  list: displays list of channels";
            qInfo() << "  quit: quits this program";
            qInfo() << "";
        }
        else
        {
            qInfo() << "Say WHAT?";
            qInfo() << "Don't understand command: " << line;
        }


    } while(true);

    return a.exec();
}
