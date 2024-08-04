#include "sendquitapp.h"

SendQuitApp::SendQuitApp(QObject *parent) : QObject{parent}
{

}

void SendQuitApp::RequestQuitApp(const QString &msg)
{
    emit PleaseQuitApp(msg);
}
