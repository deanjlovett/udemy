#ifndef TEST_H
#define TEST_H

#include <QObject>
#include <QTimer>
#include <QDebug>

class Test : public QObject
{
private:
    Q_OBJECT

    QTimer m_timer;
    uint32_t m_count;

public:
    explicit Test(QObject *parent = nullptr);

signals:

public slots:
    void timeout();

};

#endif // TEST_H
