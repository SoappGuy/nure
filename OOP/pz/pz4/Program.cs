/*
Написати статичний метод MultMatrix, який на вхід отримує двовимірний масив типу int і повертає одновимірний масив b, 
де b[i] - добуток непарних елементів, що стоять у стовпцях з від'ємним елементом в останньому рядку.
Якщо останній елемент у стовпці не від'ємний, то b[i]=-1.
Наприклад, для масиву 
{ 
    { -10,  5, -3 }, 
    {   2,  3,  5 },
    { -52, -6, 72 } 
} потрібно повернути масив {0, 15,-1}.
*/
static int[] MultMatrix(int[,] matrix)
{
    int rows = matrix.GetLength(0);
    int cols = matrix.GetLength(1);
    
    int[] result = new int[cols];
    Array.Fill(result, -1);
    
    for (int j = 0; j < cols; j++)
    {
        if (matrix[rows - 1, j] < 0)
        {
            int mult = 1;
            for (int i = 0; i < rows; i++)
            {
                  
                if (i % 2 != 0)
                {
                    mult *= matrix[i, j];
                }
            }
            result[j] = mult;
        }
    }

    return result;
}
