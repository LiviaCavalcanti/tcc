import numpy as np
import constants
import math

def calculateOraclePrice(sample):
    sorted_days = np.sort(np.unique(sample, axis=0))
    diff_access = np.diff(sorted_days) * (-1)
    
    return calculatePrice(defineS3Classes(diff_access), sorted_days)

def calculateITPrice(sample):
    days_class = numberDaysPerClass(sample)

def numberDaysPerClass(sample):
    sample = np.append(sample, constants.MAX_DAY)
    sorted_days = np.sort(np.unique(sample, axis=0))
    diff_access = np.diff(sorted_days)
    
    daysInClass = {"standard":1, "ia":0, "glacier":0, "deep_archive":0}
    for i in range(len(diff_access)):
        if diff_access[i] < 30: # é preciso contar os dias que não tiveram acesso e durante quanto tempo ele esteve numa classe 
            daysInClass["standard"] += diff_access[i] 
        else:
            classes = list(daysInClass.keys())
            number_months = math.floor(diff_access[i] / 30) # calcular o preco nas outras diff_access[i] % 30 camadas.
            number_days = diff_access[i] 

            temp = 0
            while temp <= min(number_months, len(classes) - 1):
                diff = 30 if 30 < number_days else number_days
                daysInClass[classes[temp]] += diff
                number_days -= diff
                temp += 1
            
            if(number_months > temp + 1):
                daysInClass[classes[len(classes) - 1 ]] += number_days
    return daysInClass
    
