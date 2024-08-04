#ifndef MONITOR_H
#define MONITOR_H

#include <QObject>
#include <QDebug>

class Monitor : public QObject
{
    Q_OBJECT

public:
    explicit Monitor(QObject *parent = nullptr);

signals:

public slots:
    void QuitApp(const QString& msg="");

};

#endif // MONITOR_H
