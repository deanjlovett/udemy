#include "station.h"

Station::Station(QObject *parent, int channel, QString name)
    : QObject{parent}
    , channel(channel)
    , name(name)
{
    // this->channel = channel;
    // this->name = name;
}

void Station::broadcast(QString msg)
{
    emit send(channel, name, msg);
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
