cmake_minimum_required(VERSION 3.14)

project(qt6cb-11-5-113-disconnecting LANGUAGES CXX)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(QT NAMES Qt6 Qt5 REQUIRED COMPONENTS Core)
find_package(Qt${QT_VERSION_MAJOR} REQUIRED COMPONENTS Core)

add_executable(qt6cb-11-5-113-disconnecting
  main.cpp
  radio.h radio.cpp
  station.h station.cpp
)
target_link_libraries(qt6cb-11-5-113-disconnecting Qt${QT_VERSION_MAJOR}::Core)

include(GNUInstallDirs)
install(TARGETS qt6cb-11-5-113-disconnecting
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
)
