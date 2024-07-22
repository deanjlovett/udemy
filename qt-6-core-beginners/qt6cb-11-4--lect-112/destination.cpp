#include "destination.h"
#include <QDebug> // <<== alredy in header, probably not needed
#include "qtDeansQtHelper.h"

Destination::Destination(QObject *parent)
    : QObject{parent}
{
    qInfo() << "... inside Destination::Destination(QObject *parent): QObject{parent}";
    //qInfo("... inside Destination::Destination(QObject *parent)\n\t: QObject{parent}");

    qInfo() << "...        parent " << ((parent!=nullptr) ? "!=" : "==") << " nullptr";
}

Destination::~Destination()
{
    qInfo() << "... inside Destination::~Destination()";

}

void Destination::mySignal(QString msg)
{
    qInfo() << "... inside void Destination::mySignal(QString msg)";
    //qInfo("... inside void Destination::mySignal(QString msg)");

    qInfo() << msg;

    //qStdout() << msg;
    //qStdout(msg);
}
