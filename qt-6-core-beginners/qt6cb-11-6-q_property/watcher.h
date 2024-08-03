#ifndef WATCHER_H
#define WATCHER_H

#include <QObject>
#include <QDebug>

class Watcher : public QObject
{
    Q_OBJECT

public:
    explicit Watcher(QObject *parent = nullptr);

signals:

public slots:
    void msgChanged(QString msg);

};

#endif // WATCHER_H
