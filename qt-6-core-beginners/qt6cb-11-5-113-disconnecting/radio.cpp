#include "radio.h"

Radio::Radio(QObject *parent)
    : QObject{parent}
{

}

void Radio::listen(int channel, const QString& name, const QString& msg)
{
    // qInfo() << "... inside Radio::listen(int channel: " << channel
    //         << ", QString name: " << name
    //         << ", QString msg" << msg
    //         << ")";

    qInfo() << channel << " " << name << " - " << msg;
}
