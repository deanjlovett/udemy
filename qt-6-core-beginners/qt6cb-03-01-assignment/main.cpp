#include <QCoreApplication>
#include <QDebug> // <<== might already be in QCoreApplication

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    qInfo() << "Hi, This is assignment one\n";
    qInfo() << "My name is Dean\n";


    return a.exec();
}