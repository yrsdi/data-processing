import time

start_time = time.time()

with open("/home/yrsdi/Downloads/1000000-sales-records.txt", "r") as file:
    count = 0
    for line in file:
        if count >= 1_000_000:
            break
        record = line.strip()
        # Process the record here
        print(record)
        count += 1

elapsed_time = time.time() - start_time
print("Elapsed time:", elapsed_time)
