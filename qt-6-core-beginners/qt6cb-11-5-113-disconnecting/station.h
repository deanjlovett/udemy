#ifndef STATION_H
#define STATION_H

#include <QObject>
#include <QDebug>

class Station : public QObject
{
private:
    Q_OBJECT

public:
    explicit Station(
        QObject *parent = nullptr,
        int channel = 0,
        QString name = "unknown"
    );

    // members
    QString name;
    int channel;

    // QString getName() const;
    // void setName(const QString &newName);
    QString toString();

signals:
    void send(int channel, const QString& name, const QString& msg);

public slots:
    void broadcast(const QString& msg);

};

#endif // STATION_H
