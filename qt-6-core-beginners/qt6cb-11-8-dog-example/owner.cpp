#include "owner.h"

Owner::Owner(QObject *parent)
    : QObject{parent}
{

}

void Owner::giveTreat()
{
    qInfo() << "Owner::giveTreat(), snacks are ready";
    emit treat();
}
