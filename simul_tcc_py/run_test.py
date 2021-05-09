import random
import numpy as np
import json
import math
import toolz

def generateDataUniform(low, high, sample_size):
    randSampleGet = np.random.uniform(low, high, sample_size)
    randSamplePost = np.random.uniform(low, high, sample_size)
    
    day_of_access = 1
    default = dict(zip(np.arange(1, sample_size + 1).tolist(), [[randSampleGet[i],randSamplePost[i]] for i in range(sample_size)]))
    for i in range(frequency):
        day_of_access+=random.randrange(1,sample_size)
        default[day_of_access][get] = 0
        default[day_of_access][post] = 0

def generateDataLog(totalDays, dist_type):
    if dist_type == "log":
        a = 10
        post_probability = (np.random.pareto(a, totalDays) + 1) * 1
        get_probability = np.random.lognormal(10, 10, totalDays)
        post_probability = post_probability[post_probability < 6]

        get_probability[get_probability < totalDays]
        get_probability = get_probability[get_probability < totalDays]
        return get_probability, post_probability
    
    elif dist_type == "zipf":
        a=2.1
        get_probability = np.random.zipf(a, totalDays)
        post_probability = np.random.zipf(12, totalDays)
        get_probability = get_probability[get_probability < totalDays]
    
        return get_probability, post_probability

# generates a sample of totalDays size. This sample is built using a distribution of dist_type type and with sample_size size  
def generator(totalDays, sample_size, dist_type):
    default = dict(zip(np.arange(1, sample_size + 1).tolist(), [[0,0] for i in range(totalDays)]))
    get_requests, post_requests = generateDataLog(sample_size, dist_type)
    get_requests = get_requests[get_requests < totalDays]
    
    unique, counts = np.unique(get_requests.astype(int).tolist(), return_counts=True)
    if len(post_requests) < len(counts):
            post_requests = np.append(post_requests, [[0] for i in (len(counts) - len(post_requests))])
    else:
        post_requests = post_requests[:len(counts)]
    
    default.update(dict(zip(unique.astype(int).tolist(), (np.asarray((counts.astype(int), post_requests.astype(int))).T).tolist())))

    return default

def frequencyGenerator(intervalSize, position, dictionary, randRange=20):
    day_of_access = 1
    for i in range(math.ceil(len(dictionary)/intervalSize)-1):
        day_of_access+=intervalSize

        dictionary[day_of_access][position] = random.randrange(0, randRange)
    
        # print(default)
    return dictionary

def multiplyingRequests(arr, factor=1000):
    return [ x * factor for x in arr]

if __name__ == "__main__":
    
    # number of days for the test interval
    interval_size = 1000
    # sample size of a random distribution used to build the sample
    sample_size = 10000
    # index of the type of request for every array in result `"day":[#GET, #POST]`
    get=0
    post=1

    # below is a way to distribute requests over the test interval: monthly, biannual, annual
    # represented by the number of days that the request will be considered
    # it did not look promising
    mensal=30
    biannual=180
    annual = 355
    total_rand_access = 60
    periods = [mensal, biannual, annual]
    for period in periods:
        
        default = dict(zip(np.arange(1, interval_size + 1).tolist(), [[0,0] for i in range(interval_size)]))
        get_dict = frequencyGenerator(period, get, default, randRange=20)
        for period1 in periods:
            full_dict = frequencyGenerator(period1, post, get_dict, randRange=20)

    
    # a way of produce the continuous access sample
    # another way is populate the whole sample
    a=1.3
    sample_zipf = np.random.zipf(a, 1000)
    sample_get_zipf = np.random.zipf(1.98, 1000)
    sampleDataLog = generateDataLog(10, dist_type="log")
    sampleDataLog = generator(interval_size, sample_size, dist_type="log")
    print("samples ",sampleDataLog)


    # infrequent access
    total_rand_access = 60
    interval_mx_size=5
    print(interval_size)
    default = dict(zip(np.arange(1, interval_size + 1).tolist(), [[0,0] for i in range(interval_size)]))
    for i in random.sample(range(interval_size), total_rand_access):
        access_day = i
        default[access_day][get] = random.randrange(1,500)
        default[access_day][post] = random.randrange(1,500)
        interval_random_size = random.randrange(1,interval_mx_size)
        while( interval_random_size >0 and access_day+1 < interval_size):
            default[access_day][get] = random.randrange(1,500)
            default[access_day][post] = random.randrange(1,500)
            access_day+=1
            interval_random_size-=1
    print(default)

    # if you want to multiply a previous result to have a new dimension for the same sample
    readFromFile = "your_path/here"
    with open(readFromFile) as f:
        data = json.load(f)

    sample = toolz.valmap(multiplyingRequests, data['trace'][0]['requests'])
    print(sample)
    final_dict = {"trace": [{"objSize": 10,"requests":sample}]}
    f = open("previous_file.json", "w")
    json.dump(final_dict, f)
