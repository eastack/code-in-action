#include <iostream>

int main()
{
	using namespace std;

	int secret = 10;
	int input;

	while(cin >> input)
	{
		if (input == secret){
			cout << "==" << endl;
			break;
		} else if (input > secret) {
			cout << ">" << endl;
		} else if (input < secret) {
			cout << "<" << endl;
		}
	}
}
