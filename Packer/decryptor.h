#include <Windows.h>
//-----------------------
#include <wincrypt.h>
#pragma comment(lib, "crypt32.lib")

#include <iostream>
#define AES_KEY_SIZE 16
#define CHUNK_SIZE (AES_KEY_SIZE * 3)

namespace DEC {
class Decryptor {
private:
    std::string cleartext;
    std::string password;

public:
    Decryptor(std::string pw, std::string cl);
    ~Decryptor();
    std::string decrypt();
};
}  // namespace DEC

