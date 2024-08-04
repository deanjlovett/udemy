#ifndef SENDQUITAPP_H
#define SENDQUITAPP_H

#include <QObject>

class SendQuitApp : public QObject
{
    Q_OBJECT

public:
    explicit SendQuitApp(QObject *parent = nullptr);

    void RequestQuitApp(const QString& msg = "");

signals:
    void PleaseQuitApp(const QString& msg = "");

};

#endif // SENDQUITAPP_H
