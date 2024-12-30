def merge_and_count(arr, temp, left, mid, right):
    # Initialize pointers for the left and right subarrays and the temporary array
    i, j, k = left, mid + 1, left
    inv_count = 0

    # Merge the two subarrays while counting inversions
    while i <= mid and j <= right:
        if arr[i] <= arr[j]:  # No inversion
            temp[k] = arr[i]
            i += 1
        else:  # Inversion found
            temp[k] = arr[j]
            j += 1
            inv_count += (mid - i + 1)  # Count inversions
        k += 1

    # Copy remaining elements from the left subarray (if any)
    while i <= mid:
        temp[k] = arr[i]
        i += 1
        k += 1

    # Copy remaining elements from the right subarray (if any)
    while j <= right:
        temp[k] = arr[j]
        j += 1
        k += 1

    # Copy the sorted subarray back into the original array
    for i in range(left, right + 1):
        arr[i] = temp[i]

    return inv_count

def merge_sort_and_count(arr, temp, left, right):
    # Initialize inversion count
    inv_count = 0

    if left < right:
        # Find the midpoint of the array
        mid = left + (right - left) // 2

        # Recursively sort and count inversions in the left subarray
        inv_count += merge_sort_and_count(arr, temp, left, mid)
        # Recursively sort and count inversions in the right subarray
        inv_count += merge_sort_and_count(arr, temp, mid + 1, right)
        # Merge the two sorted subarrays and count cross inversions
        inv_count += merge_and_count(arr, temp, left, mid, right)

    return inv_count

def count_inversions(arr):
    # Initialize a temporary array for merge operations
    n = len(arr)
    temp = [0] * n
    # Call the merge sort-based inversion count function
    return merge_sort_and_count(arr, temp, 0, n - 1)

def generate_array(n, m, a, b):
    # Generate an array using the provided parameters
    arr = []
    cur = 0

    for _ in range(n):
        # Update the current value based on the formula
        cur = cur * a + b
        # Calculate the array element using bit-shifting and modulo
        arr.append((cur >> 8) % m)

    return arr

def main():
    # Read input values for array size, modulo, and generation parameters
    n, m = map(int, input().split())
    a, b = map(int, input().split())

    # Generate the array based on input parameters
    arr = generate_array(n, m, a, b)

    # Debug: Print the generated array to verify correctness
    print("Generated array:", arr)

    # Compute and print the number of inversions in the array
    inversions = count_inversions(arr)
    print("Number of inversions:", inversions)

if __name__ == "__main__":
    main()
