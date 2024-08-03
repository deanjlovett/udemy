#include "watcher.h"

Watcher::Watcher(QObject *parent)
    : QObject{parent}
{}

void Watcher::msgChanged(QString msg)
{
    qInfo() << msg;

}
