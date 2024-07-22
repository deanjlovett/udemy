#ifndef DESTINATION_H // // {{
#define DESTINATION_H

#include <QObject>
#include <QDebug>

class Destination : public QObject
{
    Q_OBJECT

public:
    explicit Destination(QObject *parent = nullptr);
    ~Destination();

signals:

public slots:
    void mySignal(QString msg);

};

#endif // DESTINATION_H // }}
