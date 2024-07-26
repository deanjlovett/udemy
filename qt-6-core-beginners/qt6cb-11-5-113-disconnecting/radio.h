#ifndef RADIO_H
#define RADIO_H

#include <QObject>
#include <QDebug>

class Radio : public QObject
{
private:
    Q_OBJECT

public:
    explicit Radio(QObject *parent = nullptr);

signals:
    void quit();


public slots:
    void listen(int channel, const QString& name, const QString& msg);


};

#endif // RADIO_H
