namespace List;
// 15. 
// 1✓, 2✓, 6✓, 7✓, 15✓, 19✓, 27✓, 41✓, 61✓, 77✓, 79✓, 54(5)-


public class Node
{
    public int Data;
    public Node? Next;
    
    public Node(int data)
    {
        this.Data = data;
        this.Next = null;
    }
    public Node(int data, Node? next)
    {
        this.Data = data;
        this.Next = next;
    }

    public void PushRecursive(int data)
    {
        if ( this.Next == null)
        {
            this.Next = new Node(data);
        }
        else
        {
            this.Next.PushRecursive(data);
        }
    }
    public int PopRecursive()
    {
        int data;
        if ( this.Next.Next == null)
        {
            data = this.Next.Data;
            this.Next = null;
        }
        else
        {
            data = this.Next.PopRecursive();
        }

        return data;
    }
    public override string ToString()
    {
        return this.Data.ToString();
    }
}

public class List
{
    private Node? Head { get; set; }
    public int Length { get; private set; }

    public bool IsShort
    {
        get
        {
            return this.Length <= 5;
        }

        set
        {   
            if (this.Head == null)
            {
                Console.WriteLine("list is empty");
                return;
            }
            
            if (value)
            {
                Node currNode = this.Head;
                for (int i = 0; i < (Math.Min(4, this.Length - 1)); i++)
                {
                    currNode = currNode.Next;
                }

                currNode.Next = null;
                this.Length = 5;
            }
            else
            {
                this.PushRecursive(5);
            }
        }
    }
    public int this[int start, int end]
    {
        get
        {
            if (this.Head == null)
            {
                Console.WriteLine("list is empty");
                return -1;
            }
            if (start < 0)
            {
                start = 0;
            }
            if (end >= this.Length)
            {
                end = this.Length - 1;
            }
            if (start >= this.Length )
            {
                start = this.Length - 1;
            }
            if (end < 0)
            {
                end = 0;
            }
            if (start > end)
            {
                (start, end) = (end, start);
            }
            
            Node max = this.Head;
            for (int i = 0; i < start; i++)
            {
                max = max.Next;
            }

            Node currNode = max;
            for (int i = start; i < end; i++)
            {
                currNode = currNode.Next;
                if (currNode.Data > max.Data)
                {
                    max = currNode;
                }
                
            }
            return max.Data;
        }
    }
    
    public List()
    {
        this.Length = 0;
        this.Head = null;
    }
    public List(int data)
    {
        this.Length = 1;
        this.Head = new Node(data);
    }
    public List(int data, Node? nextNode)
    {
        this.Length = 1;
        this.Head = new Node(data, nextNode);
    }
    
    public void PushRecursive(int data)
    {
        this.Length++;
        
        if (this.Head == null)
        {
            this.Head = new Node(data);
        }
        else
        {
            this.Head.PushRecursive(data);
        }
        
    }
    public int PopRecursive()
    {
        if (this.Head == null)
        {
            throw new InvalidOperationException("Cannot pop from an empty list.");
        }

        this.Length--;
        int data;
        if (this.Head.Next == null)
        {
            data = this.Head.Data;
            this.Head = null;
        }
        else
        {
            data = this.Head.PopRecursive();
        }

        return data;

    }
    public void Insert(int data, int index)
    {
        if (index >= this.Length)
        {
            index = this.Length;
        }

        this.Length++;
        if (index <= 0)
        {
            this.Head = new Node(data, this.Head);
        }
        
        else if (index > 0)
        {
            var currNode = this.Head;
            for (int i = 0; i < index - 1 && currNode.Next != null; i++)
            {
                currNode = currNode.Next;
            }

            currNode.Next = new Node(data, currNode.Next);
            
        }
    }
    public int FilteredDelete(int data)
    {
        if (this.Head == null)
        {
            Console.WriteLine("Cannot delete from an empty list.");
            return 0;
        }
        
        int count = 0;
        var currNode = this.Head;
        for (int i = 0; i < this.Length - 1; i++)
        {
            while (currNode.Next != null && currNode.Next.Data == data)
            {
                count++;
                this.Length--;
                currNode.Next = currNode.Next.Next;
            }

            currNode = currNode.Next;
        }

        if (this.Head.Data == data)
        {
            count++;
            this.Head = this.Head.Next;
        }

        return count;
        
    }
    public void ColumnPrint()
    {
        Console.WriteLine("-----");
        if (this.Head == null)
        {
            Console.WriteLine("Empty");
        }
        
        var currNode = this.Head;
        for (int i = 0; i < this.Length; i++)
        {
            Console.WriteLine(currNode);
            currNode = currNode.Next;
        }
        Console.WriteLine("-----");
        
    }
    public int Search(int data)
    {
        if (this.Head == null)
        {
            throw new InvalidOperationException("Cannot search in an empty list.");
        }
        
        var currNode = this.Head;
        for (int i = 0; i < this.Length - 1; i++)
        {
            if (currNode.Data == data)
            {
                return i;
            }

            currNode = currNode.Next;
        }
        return -1;
        
    }
    public void Empty()
    {
        this.Length = 0;
        this.Head = null;
    }
    
    public static implicit operator bool(List lst)
    {
        if (lst.Length > 0)
        {
            return true;
        }
        else
        {
            return false;
        }
    }

    public static List operator ++(List lst)
    {
        if (lst.Head == null)
        {
            Console.WriteLine("list is empty");
            return lst;
        }
        Node currNode = lst.Head;
        for (int i = 0; i < lst.Length; i++)
        {
            currNode.Data++;
            currNode = currNode.Next;
        }

        return lst;
    }
    public override string ToString()
    {
        string repr = "";
        
        Node? currNode = this.Head;
        while (currNode != null)
        {
            repr += currNode + " -> ";
            currNode = currNode.Next;
        }

        if (repr.Length == 0)
        {
            return "Empty";
        }
        
        return repr[..^4];
    }
}