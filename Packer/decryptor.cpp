#include "decryptor.h"

DEC::Decryptor::Decryptor(std::string pw, std::string cl)
    : password(
          "PRz5trEwNgA52wPZbR4GgOW1k9SC1ibiVF8Sio3SHQRtSOPc9TiWk89uMKKApbOzVmQC"
          "sTsK8vY8V0WAMnWfwEO72bZ1ZvVn51C")
{
    cleartext = cl;
    password = pw;
}

int main() {}

