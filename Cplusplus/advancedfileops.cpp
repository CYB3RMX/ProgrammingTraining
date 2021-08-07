#include <iostream>
#include <fstream>
#include <string>
#include <cstring>

using namespace std;

class FileOps{
    public:
        void OpenFile(string filename){
            // Buffer
            string text;

            // Open a file in read mode
            ifstream file_handler(filename);
            while(getline(file_handler, text)){
                cout << text << endl;
            }
            file_handler.close();
        }
        void WriteFile(string filename, string buffer){
            // Open a file in write mode
            ofstream file_handler(filename);
            if(buffer.size() > 100){ // Check for buffer size
                cout << "High amount of buffer size detected!!" << endl;
                exit(1);
            }
            else{
                file_handler << buffer;
            }
            file_handler.close();
        }
};

int main(){
    // Creating a fileops object
    FileOps fpos;

    // Target file names to read and write operations
    string file_read;
    string file_write;

    // Buffer for write
    char writebuff[100];

    // Read from file
    cout << "Enter file to read: ";
    cin >> file_read;
    fpos.OpenFile(file_read);

    // Write to file
    cout << "Enter buffer to write: ";
    cin >> writebuff;
    fpos.WriteFile("saved.txt", writebuff);
    cout << "\nBuffer saved into saved.txt" << endl;

    return 0;
}