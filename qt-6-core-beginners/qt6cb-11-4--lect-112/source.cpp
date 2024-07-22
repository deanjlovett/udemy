#include "source.h"
#include <QDebug> // <<== alredy in header, probably not needed
#include "qtDeansQtHelper.h"

Source::Source(QObject *parent)
    : QObject{parent}
{
    qInfo() << "... inside Source::Source(QObject *parent): QObject{parent}";

    qInfo() << "...        parent " << ((parent !=nullptr) ? "!=" : "==") << " nullptr";

}

void Source::test()
{
    qInfo() << "... inside void Source::test()";

    emit mySignal("Hello World!");
}
