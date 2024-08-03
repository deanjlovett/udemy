/*
    What:
    Qt Property System

    Qt Property Binding
    Qt Property Bindings

    QML
*/

#include <QCoreApplication>
#include "test.h"
#include "watcher.h"

#include <QVariant>


int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    // Set up code that uses the Qt event loop here.
    // Call a.quit() or a.exit() to quit the application.
    // A not very useful example would be including
    // #include <QTimer>
    // near the top of the file and calling
    // QTimer::singleShot(5000, &a, &QCoreApplication::quit);
    // which quits the application after 5 seconds.

    // If you do not need a running Qt event loop, remove the call
    // to a.exec() or use the Non-Qt Plain C++ Application template.

    Test tester;
    Watcher watcher;

    QObject::connect(
        &tester,
        &Test::msgChanged,     // signal
        &watcher,
        &Watcher::msgChanged   // slot
    );

    tester.setProperty("msg",QVariant("Hello World #1, via setProperty"));
    tester.setMsg(                    "Hello World #2, via setMsg"      );


    return a.exec();
}
