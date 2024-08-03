#include "test.h"

Test::Test(QObject *parent)
    : QObject{parent}
    , m_count(0)
{
    this->connect(
        &m_timer,
        &QTimer::timeout,
        this,
        &Test::timeout
    );
    m_timer.setInterval(1000);
    m_timer.start();
    // m_timer.stop();

    // m_count = 0; // moved up

}

void Test::timeout()
{

    qInfo() << "Test: timeout called. Count: " << ++m_count;

    if( m_count >= 4 ) {
        m_timer.stop();
    }
}
