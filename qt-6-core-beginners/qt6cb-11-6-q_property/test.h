#ifndef TEST_H
#define TEST_H

#include <QObject>

class Test : public QObject
{
    Q_OBJECT

    QString m_msg;

public:
    explicit Test(QObject *parent = nullptr);

    // Q_PROPERTY(type name READ name WRITE setName NOTIFY nameChanged FINAL)
    //
    Q_PROPERTY(QString msg READ msg WRITE setMsg NOTIFY msgChanged FINAL)

    QString msg() const;
    void setMsg(const QString &newMsg);

signals:

    void msgChanged(QString msg);

};

#endif // TEST_H
