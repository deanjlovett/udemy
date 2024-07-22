#ifndef QTDEANSQTHELPER_H
#define QTDEANSQTHELPER_H

#include <QTextStream>

inline QTextStream& qStdout()
{
    static QTextStream r{stdout};
    return r;
}

inline QTextStream& qStdout(QString& qstr)
{
    static QTextStream qts{stdout};
    qts << qstr;
    return qts;
}



#endif // QTDEANSQTHELPER_H
