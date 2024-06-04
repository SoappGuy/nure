// Написати метод, вхідними параметрами якого є рядок, символ. Метод повинен повернути номер позиції заданого символу в рядку, якщо він там є, і 0, якщо його там немає.
// Задача. Написати метод, який нічого не повертає, отримує на вхід масив цілих чисел та видаляє з нього парні значення.
// Задача. Написати метод, який нічого не повертає, отримує на вхід масив цілих чисел та видаляє з нього парні значення.

namespace pz2;

class Program
{

    static void Task1()
    {   
        // Заданий масив. Створити новий масив, в якому залишити лише унікальні елементи першого,
        // наприклад, [1,3,2,4,3,2,1] => [1,3,2,4]. 
        int[] arr = new int[] {1, 3, 2, 4, 3, 2, 1};
        int[] hashed = new int[arr.Length];
        int len = 0;
        
        foreach (int element in arr)
        {
            if (hashed.Contains(element)) continue;

            hashed[len] = element;
            len++;
        }
        
        Console.WriteLine(string.Join(", ", hashed[..len]));
    }

    static void Task2()
    {
        // Змініть програму так, щоб при вході в неї у користувача питалося,
        // суму чого порахувати (стовпчика або рядка), питався номер,
        // і порахувати суму елементів відповідного стовпця або відповідного рядка.
        Console.WriteLine("1 - row\n2 - column");
        int question = Console.Read();

    }
    
    static void Main(string[] args)
    {
        var DoMath = (double R, double H, double L) => {
            double A = R;
            double B = R - H/2;
            double C = L;
            
            double cosA = ((A * A) + (B * B) - (C * C) ) / (2 * A * B);
            return Math.Acos(cosA);
        };
        
        double R = 10;
        double R1 = 10;
        double R2 = 10;
        double H = 10;
        double X = 10;
        double Y = 10;

        double topAngle = 2 * DoMath(R, H, R1);
        double botAngle = 2 * DoMath(R, H, R2);

        double edgeAngle = 360 - topAngle - botAngle;

        double start1 = edgeAngle / 2 + botAngle;
        double start2 = start1 + edgeAngle + topAngle;

        var topLine = (
            ((X + R - R1), (Y + R - (H/2))),
            ((X + R + R1), (Y + R - (H/2)))
            );
        var botLine = (
            ((X + R - R2), (Y + R + (H/2))),
            ((X + R + R2), (Y + R + (H/2)))
        );

        var topArc = (
            (X, Y, (R + R), (R + R)), 
            start1, edgeAngle
            );
        var botArc = (
            (X, Y, (R + R), (R + R)), 
            start2, edgeAngle
        );

        path.Data = Geometry.Parse(
            $"M{Math.Round(m2x)},{Math.Round(my)} " +
            $"A{originalRadius},{originalRadius} 0 0 1 {Math.Round(l2x)},{Math.Round(ly)} " +
            $"L{Math.Round(l1x)},{Math.Round(ly)} " +
            $"A{originalRadius},{originalRadius} 0 0 1 {Math.Round(m1x)},{Math.Round(my)} " +
            $"z");

        float botLeftX = X;
        float topLeftX = X;

        $"M{},{} " +
        $"A{originalRadius},{originalRadius} 0 0 1 {Math.Round(l2x)},{Math.Round(ly)} " +
        $"L{Math.Round(l1x)},{Math.Round(ly)} " +
        $"A{originalRadius},{originalRadius} 0 0 1 {Math.Round(m1x)},{Math.Round(my)} " +
        $"z"
    }
}