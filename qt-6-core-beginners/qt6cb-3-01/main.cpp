/*
    What:
    Hello world

    Why:
    This is the starting point in just about every language.
    Teaches you what a basic application looks like.

    How:
    Main Function
    Includes
    QCoreApplication and exec
    CMake
    CMake Modules


*/

#include <QCoreApplication>
#include <QDebug>

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    qInfo() << "Hello world!" << "\n";

    return a.exec();
}
