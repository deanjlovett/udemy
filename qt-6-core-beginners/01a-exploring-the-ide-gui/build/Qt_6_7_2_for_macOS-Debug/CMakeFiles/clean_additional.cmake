# Additional clean files
cmake_minimum_required(VERSION 3.16)

if("${CONFIG}" STREQUAL "" OR "${CONFIG}" STREQUAL "Debug")
  file(REMOVE_RECURSE
  "01a-exploring-the-ide-gui_autogen"
  "CMakeFiles/01a-exploring-the-ide-gui_autogen.dir/AutogenUsed.txt"
  "CMakeFiles/01a-exploring-the-ide-gui_autogen.dir/ParseCache.txt"
  )
endif()
