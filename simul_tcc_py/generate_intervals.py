import itertools
import json

interval_size=5

#
def get_index_array(i, j, arranged_seq):
    return_index = -1
    return_first = -1
    return_last = -1
    for index in range(len(arranged_seq)):
        first,last = arranged_seq[index][-1]
        if last == i and first ==j:
            
            return_index=index
            return_first = first 
            return_last = last
            break

    return [(return_first, return_last)], return_index

# concat the updgraded one array of intervals with the a
def concat_array(tt, sec,temp_index):
    to_substitute = sec[temp_index][:-1]
    new_tt_sec = []
    for elem in tt:
        new_tt_sec.append(to_substitute+elem)
    if temp_index == 0:
        return new_tt_sec + sec[1:]
    elif temp_index == len(sec) - 1:
        return sec[-1] + new_tt_sec
    else:
        return sec[:temp_index] + new_tt_sec + sec[temp_index+1:]

# eliminate redundancy in the intervals of an array
def eliminate_first(arr):
    arr_return = [[arr[0][0]]]
    for i in range(len(arr)):
        if arr[i][0] == arr[i][1] -1:
            arr_return.append([arr[i][1]])
        else:
            if len(arr_return) >= 1:
                if len(arr_return[-1]) == 1 and arr_return[-1][0] ==arr[i][0]:
                    arr_return = arr_return[:-1]
                elif len(arr_return[-1]) == 2 and arr_return[-1][1] ==arr[i][0]:
                    arr_return.append([arr[i][0] + 1, arr[i][1]])
                    continue
            arr_return.append([arr[i][0], arr[i][1]])
    return arr_return

def eliminate_last(arr):
    arr_return = []
    for i in range(len(arr)):
        if arr[i][0] == arr[i][1] -1:
            if len(arr_return) >= 1:
                
                if arr_return[-1][-1] == arr[i][0]:
                    arr_return.append([arr[i][1]])
                    continue

            arr_return.append([arr[i][0]])
        else:
            if len(arr_return) >= 1:
                
                if len(arr_return[-1]) == 1 and arr_return[-1][0] ==arr[i][0]:
                    arr_return = arr_return[:-1]
                if len(arr_return[-1]) == 2 and arr_return[-1][1] ==arr[i][0]:
                    arr_return[-1][1] -= 1
                    
            arr_return.append([arr[i][0], arr[i][1]])

    return arr_return
 
basic_seq = list(range(1, interval_size+1))
arranged_seq=list(itertools.combinations(basic_seq, 2))
sec = []
for e in arranged_seq:
    if e[0] ==1:
        sec.append([e])

for j in range(1,interval_size):
    for i in range(j+1, interval_size):

        temp_array, temp_index = get_index_array(i, j, sec)
        tt = []
        for arr in temp_array:
            for (first,last) in arranged_seq:
                if first == arr[1]:
                    tt.append([arr, (first, last)])
                    
        
        sec = concat_array(tt, sec,temp_index)

fileName = "intervals_array1.json"
f = open(fileName, "w")
json.dump(sec, f)
f.close()
final_answer = []

for elem in sec:
    new_elem = eliminate_first(elem)
    if new_elem[-1][-1] != interval_size:
        new_elem.append([interval_size])
    final_answer.append(new_elem)

    new_elem = eliminate_last(elem)
    if new_elem[-1][-1] != interval_size:
        new_elem.append([interval_size])
    final_answer.append(new_elem)
    
fileName = "intervals_array2.json"
f = open(fileName, "w")
json.dump(final_answer, f)
f.close()
