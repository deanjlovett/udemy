#include <string>
#include <sstream>
#include "station.h"

Station::Station(QObject *parent, int channel, QString name)
    : QObject{parent}
    , channel(channel)
    , name(name)
{
    // this->channel = channel;
    // this->name = name;
}

void Station::broadcast(const QString& msg)
{
    emit send(channel, name, msg);
}

QString Station::toString()
{
    QString ss;
    ss.append( std::to_string(this->channel) );
    ss.append(" ");
    ss.append(this->name);

    return ss;
}

// void send(int channel, const QString& name, const QString& msg)
// {
//
// }


// QString Station::getName() const
// {
//     return name;
// }

// void Station::setName(const QString &newName)
// {
//     name = newName;
// }
