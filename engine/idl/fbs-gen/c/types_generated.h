// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_TYPES_H_
#define FLATBUFFERS_GENERATED_TYPES_H_

#include "flatbuffers/flatbuffers.h"

enum DataType {
  INT = 0,
  LONG = 1,
  FLOAT = 2,
  DOUBLE = 3,
  STRING = 4,
  VECTOR = 5
};

inline const DataType (&EnumValuesDataType())[6] {
  static const DataType values[] = {
    INT,
    LONG,
    FLOAT,
    DOUBLE,
    STRING,
    VECTOR
  };
  return values;
}

inline const char * const *EnumNamesDataType() {
  static const char * const names[] = {
    "INT",
    "LONG",
    "FLOAT",
    "DOUBLE",
    "STRING",
    "VECTOR",
    nullptr
  };
  return names;
}

inline const char *EnumNameDataType(DataType e) {
  if (e < INT || e > VECTOR) return "";
  const size_t index = static_cast<size_t>(e);
  return EnumNamesDataType()[index];
}

#endif  // FLATBUFFERS_GENERATED_TYPES_H_
