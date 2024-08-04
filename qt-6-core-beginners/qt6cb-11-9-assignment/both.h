#ifndef BOTH_H
#define BOTH_H

#include <QObject>

class Both : public QObject
{
    Q_OBJECT
public:
    explicit Both(QObject *parent = nullptr);

signals:
};

#endif // BOTH_H
