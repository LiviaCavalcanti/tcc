classesPrice = [{"name":"archive", "storagePrice": 0.00099, "getPrice": 0.0000004, "postPrice": 0.00005, "retrievalPrice": [0.00010, 0.000025], "retrievalData": [0, 0.02, 0.0025], "days": 180},
	{"name":"glacier", storagePrice: [0.004], "getPrice": 0.0000004, "postPrice": 0.00005, "retrievalPrice": [0.01000, 0.00005, 0.000025], "retrievalData": [0.03, 0.01, 0.0025], "days": 90},
	{"name":"infrequent access", "storagePrice": [0.0125], "getPrice": 0.001, "postPrice": 0.01, "retrievalData": [0.01], "days": 30},
	{"name":"standard", "storagePrice": [0.023, 0.022, 0.021], "getPrice": 0.0000004, "postPrice": 0.000005, "days": 30}]

get_request = list(range(1, 1000000, 20))
post_request = list(range(1, 1000000, 20))

def calculatePrice(classPosition, get_request, post_request):
    get_cost = classesPrice[classPosition]["get_request"] * get_request + classesPrice[classPosition]["post_request"] *post_request 

Z = calculatePrice(3, get_request, post_request)
fig = plt.figure()
ax = plt.axes(projection='3d')
ax.contour3D(get_request, post_request, Z, 50, cmap='binary')
