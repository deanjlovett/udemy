/*
 *
 *
 */

#include <QCoreApplication>
#include "source.h"
#include "destination.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    Source oSrc(&a);
    Destination oDst(&a);

    // Source oSrc;
    // Destination oDst;


    QObject::connect(
        &oSrc,
        &Source::mySignal,
        &oDst,
        &Destination::mySignal
    );

    oSrc.test();

    oDst.mySignal("direct call.");


    return a.exec();
}
