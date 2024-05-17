namespace List;

class Program
{
    static void Main(string[] args)
    {
        List lst = new List();
        lst.PushRecursive(10);
        lst.PushRecursive(10);
        lst.PushRecursive(115);
        Console.WriteLine(lst.ToString());
        lst.IsShort = false;
        // Console.WriteLine(lst.FilteredDelete(115));
        Console.WriteLine(lst.ToString());

        // // PushRecursive
        // Console.WriteLine("\nPushRecursive: ");
        // lst.PushRecursive(3);
        // lst.PushRecursive(7);
        // Console.WriteLine(lst.ToString());
        //
        // // Insert
        // Console.WriteLine("\nInsert: ");
        // lst.Insert(3, 0);
        // lst.Insert(4, 2);
        // Console.WriteLine(lst.ToString());
        //
        // // PopRecursive
        // Console.WriteLine("\nPopRecursive: ");
        // Console.WriteLine(lst.PopRecursive());
        // Console.WriteLine(lst.PopRecursive());
        // Console.WriteLine(lst.ToString());
        //
        // // FilteredDelete
        // Console.WriteLine("\nFilteredDelete: ");
        // Console.WriteLine(lst.FilteredDelete(3));
        // Console.WriteLine(lst.ToString());
        //
        // // ColumnPrint
        // Console.WriteLine("\nColumnPrint: ");
        // lst.ColumnPrint();
        // lst.PushRecursive(2);
        // lst.PushRecursive(7);
        // lst.ColumnPrint();
        //
        // // Search
        // Console.WriteLine("\nSearch: ");
        // for (int i = 0; i < 5; i++)
        // {
        //     lst.PushRecursive(Random.Shared.Next() % 10);
        // }
        // Console.WriteLine(lst.ToString());
        // Console.WriteLine(lst.Search(8));
        //
        // // Max
        // Console.WriteLine("\nMax: ");
        // Console.WriteLine(lst[2, 5]);
        //
        // // bool implicit operator
        // Console.WriteLine("\nbool implicit operator: ");
        // Console.WriteLine(lst);
        // lst.Empty();
        // Console.WriteLine(lst);
    }
}