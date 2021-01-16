#include <Windows.h>
//-------------------------
#include <wincrypt.h>
#pragma comment(lib, "crypt32.lib")

#include <iostream>

#include "strings.h"

#ifndef __encryptor_h
#define __encryptor_h

#define AES_KEY_SIZE 16
#define CHUNK_SIZE (AES_KEY_SIZE * 3)

namespace ENC {
class Encryptor {
private:
    std::string cleartext;
    std::string password;

public:
    Encryptor(std::string pw, std::string cl);
    ~Encryptor();
    std::string encrypt();
};
}  // namespace ENC

#endif
