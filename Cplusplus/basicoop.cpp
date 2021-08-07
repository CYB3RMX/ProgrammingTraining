#include <iostream>

using namespace std;

class Box{
    public:
        int length;
        int height;
        int breadth;
};

int main(){
    Box box1;
    box1.height = 10;
    box1.length = 10;
    box1.breadth = 20;

    cout << "All: " << box1.height * box1.length * box1.breadth << endl;
    return 0;
}