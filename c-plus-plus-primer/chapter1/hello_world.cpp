#include <iostream>

/*
 * print "Hello World!"
 */
int main() 
{
	std::cout << "Hello " << std::flush;

	// use std namespace
	using namespace std;
	cout << "World!" << std::endl;

	return 0;
}
