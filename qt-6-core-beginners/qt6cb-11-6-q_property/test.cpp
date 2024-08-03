#include "test.h"


Test::Test(QObject *parent) : QObject{parent}
{

}



QString Test::msg() const
{
    return m_msg;
}

void Test::setMsg(const QString &newMsg)
{
    m_msg = newMsg;
    emit msgChanged(m_msg);
}
