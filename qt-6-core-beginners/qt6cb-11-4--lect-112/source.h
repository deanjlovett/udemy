#ifndef SOURCE_H
#define SOURCE_H

#include <QObject>
#include <QDebug>

class Source : public QObject
{
    Q_OBJECT

public:
    explicit Source(QObject *parent = nullptr);

    void test();

signals:
    void mySignal(QString msg);

};

#endif // SOURCE_H
