#include <iostream>

using namespace std;

class Abilities{
    public:
        void Run(){
            cout << "Running.." << endl;
        }
        void Stop(){
            cout << "Stop!!" << endl;
        }
};

class Human: public Abilities{
    public:
        int age;
};

int main(){
    Human hum;
    hum.Run();
    hum.Stop();
    return 0;
}