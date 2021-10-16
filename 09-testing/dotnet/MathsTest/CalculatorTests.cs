using dotnet.Maths;
using Xunit;

namespace dotnet.MathsTest
{
    public class CalculatorTests
    {
        [Fact]
        public void Add_operation_should_return_correct_value()
        {
            //Given
            var calculator = new Calculator();
            var a = 10;
            var b = 2;
            var expected = 12;

            //When
            var actual = calculator.Add(a, b);

            //Then
            Assert.Equal(expected, actual);
        }

        [Fact]
        public void Subtract_operation_should_return_correct_value()
        {
            //Given
            var calculator = new Calculator();
            var a = 10;
            var b = 2;
            var expected = 8;

            //When
            var actual = calculator.Subtract(a, b);

            //Then
            Assert.Equal(expected, actual);
        }

        [Fact]
        public void Multiply_operation_should_return_correct_value()
        {
            //Given
            var calculator = new Calculator();
            var a = 10;
            var b = 2;
            var expected = 20;

            //When
            var actual = calculator.Multiply(a, b);

            //Then
            Assert.Equal(expected, actual);
        }

        [Fact]
        public void Divide_operation_should_return_correct_value()
        {
            //Given
            var calculator = new Calculator();
            var a = 10;
            var b = 2;
            var expected = 5;

            //When
            var actual = calculator.Divide(a, b);

            //Then
            Assert.Equal(expected, actual);
        }
    }
}
