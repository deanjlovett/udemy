#include "monitor.h"

Monitor::Monitor(QObject *parent)
    : QObject{parent}
{

}

void Monitor::QuitApp(const QString &msg)
{
    qInfo() << "";
    qInfo() << "... inside Monitor::QuitApp(const QString &msg: """
            << ((msg=="") ? "((empty))" : msg)
            << """";
    qInfo() << "... msg: " << msg;
    qInfo() << "";
}
