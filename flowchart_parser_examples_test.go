package visibleIdeas

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_flowchartJsonStringToWorkflow_example_02(t *testing.T) {
	RegisterTestingT(t)

	{

		s := `{"nodes":[{"id":"sensorContainer","type":"ContainerNode","data":{"label":"Signs & Symptoms","energized":""},"position":{"x":100,"y":0},"className":"bg-blue-300/50 rounded-lg border-0 ring-2","style":{"width":400,"height":5000,"pointerEvents":"none"},"draggable":false,"positionAbsolute":{"x":100,"y":0},"z":0,"handleBounds":{"source":null,"target":null},"width":400,"height":5000,"isParent":true},{"id":"findingContainer","type":"ContainerNode","data":{"label":"Findings","energized":""},"position":{"x":575,"y":0},"className":"bg-emerald-300/50 rounded-lg border-0 ring-emerald-400 ring-2","style":{"width":400,"height":5000,"pointerEvents":"none"},"draggable":false,"positionAbsolute":{"x":575,"y":0},"z":0,"handleBounds":{"source":null,"target":null},"width":400,"height":5000,"isParent":true},{"id":"logicContainer","type":"ContainerNode","data":{"label":"Logic","energized":""},"position":{"x":1050,"y":0},"className":"bg-zinc-300/50 rounded-lg border-0 ring-emerald-400 ring-2","style":{"width":400,"height":5000,"pointerEvents":"none"},"draggable":false,"positionAbsolute":{"x":1050,"y":0},"z":0,"handleBounds":{"source":null,"target":null},"width":400,"height":5000,"isParent":true},{"id":"ddxContainer","type":"ContainerNode","data":{"label":"DDx","energized":""},"position":{"x":1525,"y":0},"className":"bg-purple-300/50 rounded-lg border-0 ring-purple-400 ring-2","style":{"width":400,"height":5000,"pointerEvents":"none"},"draggable":false,"positionAbsolute":{"x":1525,"y":0},"z":0,"handleBounds":{"source":null,"target":null},"width":400,"height":5000,"isParent":true},{"id":"hiJqpvH5hk48vIEZ","type":"SensorNode","position":{"x":17,"y":953.75},"data":{"nodeID":"hiJqpvH5hk48vIEZ","identifier":{"id":"SENS-Target-Cell-count-over-RBC-count","source":"sour-candy"},"label":"Target cell count over RBC count","short":"TRG / RBC","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":117,"y":953.75},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":112.74993896484375,"y":20,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.00018310546875,"y":20,"width":16,"height":16}]},"width":125,"height":56,"selected":false,"dragging":false},{"id":"TLbDaPOs9CO4DbA7","type":"FindingComputeNode","position":{"x":43,"y":922.75},"data":{"nodeID":"TLbDaPOs9CO4DbA7","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Target Cells detected","functionality":"reported","params":[{"value":0}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":618,"y":922.75},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":187.9998779296875,"y":17.9998779296875,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.00018310546875,"y":17.9998779296875,"width":16,"height":16}]},"width":265,"height":72,"selected":false,"dragging":false},{"id":"zVHcsLQ1twf5cv0K","type":"SensorNode","position":{"x":21,"y":1035.3133012253477},"data":{"nodeID":"zVHcsLQ1twf5cv0K","identifier":{"id":"SENS-Hemoglobin-concentration","source":"sour-candy"},"label":"Hemoglobin-concentration","short":"Hb","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":121,"y":1035.3133012253477},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":57.09722572652425,"y":19.994731268384538,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.016350116801863,"y":19.994731268384538,"width":16,"height":16}]},"width":69,"height":56,"selected":false,"dragging":false},{"id":"EUnVZT4c9Ub19d0B","type":"SensorNode","position":{"x":10,"y":1164.569165565408},"data":{"nodeID":"EUnVZT4c9Ub19d0B","identifier":{"id":"SENS-Mean-Corpuscular-Volume","source":"sour-candy"},"label":"Mean-Corpuscular-Volume","short":"MCV","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":110,"y":1164.569165565408},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":71.31773964311247,"y":19.994731268384538,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.016350116801863,"y":19.994731268384538,"width":16,"height":16}]},"width":83,"height":56},{"id":"EfLwQJuY3Heo1bpg","type":"FindingComputeNode","position":{"x":52,"y":1055.7416664414864},"data":{"nodeID":"EfLwQJuY3Heo1bpg","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Low Hemoglobin","functionality":"less_than","params":[{"value":"11","units":{"id":"12016","name":"grams per deciliter","short_1":"g/dl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":627,"y":1055.7416664414864},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":220.43390943739942,"y":27.980318702772983,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.9855874343884485,"y":27.980318702772983,"width":16,"height":16}]},"width":232,"height":72,"selected":false,"dragging":false},{"id":"33DQTv4BkYYmwBQu","type":"FindingComputeNode","position":{"x":52,"y":1159.7402229544577},"data":{"nodeID":"33DQTv4BkYYmwBQu","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Low MCV","functionality":"less_than","params":[{"value":"80","units":{"id":"13023","name":"femtoliter","short_1":"fl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":627,"y":1159.7402229544577},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":188.00003164812122,"y":17.998362529241184,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.9855874343884485,"y":17.998362529241184,"width":16,"height":16}]},"width":200,"height":72,"selected":false,"dragging":false},{"id":"NJmkKTJBqxaInVup","type":"FindingComputeNode","position":{"x":52,"y":1299.9975307815466},"data":{"nodeID":"NJmkKTJBqxaInVup","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Normal or High Hemoglobin","functionality":"greater_than_or_equal","params":[{"value":"11","units":{"id":"12016","name":"grams per deciliter","short_1":"g/dl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":627,"y":1299.9975307815466},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":295.2530272237968,"y":28.01102514627888,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.985699912203488,"y":28.01102514627888,"width":16,"height":16}]},"width":307,"height":72,"selected":false,"dragging":false},{"id":"e5q43Fpckvkz36UN","type":"FindingComputeNode","position":{"x":44,"y":1387.866711637459},"data":{"nodeID":"e5q43Fpckvkz36UN","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Normal or High MCV","functionality":"greater_than_or_equal","params":[{"value":"80","units":{"id":"13023","name":"femtoliter","short_1":"fl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":619,"y":1387.866711637459},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":188.00003164812122,"y":17.998250051426144,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.9855874343884485,"y":17.998250051426144,"width":16,"height":16}]},"width":258,"height":72,"selected":false,"dragging":false},{"id":"2OP30MUc57jrcwKX","type":"LogicNode","position":{"x":10,"y":1090.1415820591974},"data":{"nodeID":"2OP30MUc57jrcwKX","identifier":{"id":"orNode","source":"sour-candy"},"label":"OR","weight":0},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1060,"y":1090.1415820591974},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":67.99992820763377,"y":31.982726250545674,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.995748721069461,"y":31.982726250545674,"width":16,"height":16}]},"width":80,"height":80,"zIndex":1},{"id":"fzMlRxUGtB3odOX2","type":"LogicNode","position":{"x":10,"y":1378.1576063344387},"data":{"nodeID":"fzMlRxUGtB3odOX2","identifier":{"id":"orNode","source":"sour-candy"},"label":"OR","weight":0},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1060,"y":1378.1576063344387},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":68.01487475334496,"y":32.0170756622737,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.995754771980186,"y":32.0170756622737,"width":16,"height":16}]},"width":80,"height":80,"zIndex":1},{"id":"1ePdHT3I06tb06zw","type":"DDxNode","position":{"x":15,"y":934.9754372371881},"data":{"nodeID":"1ePdHT3I06tb06zw","identifier":{"id":"DDX-Iron-Deficiency-Anemia","source":"sour-candy"},"label":"Iron Deficiency Anemia","short":"Iron Deficiency Anemia","weight":0},"parentNode":"ddxContainer","extent":"parent","positionAbsolute":{"x":1540,"y":934.9754372371881},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":203.35834030231723,"y":32.0170756622737,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.995629101964124,"y":32.0170756622737,"width":16,"height":16}]},"width":215,"height":80,"selected":false,"dragging":false},{"id":"9DYkQJccdNqQSSLy","type":"WeightNode","position":{"x":272,"y":1057.5933617863825},"data":{"nodeID":"9DYkQJccdNqQSSLy","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":1},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1322,"y":1057.5933617863825},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":84.00073230060394,"y":39.9914812501567,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.014790165003332,"y":39.9914812501567,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"5HWIYkSdp12TH7pt","type":"SensorNode","position":{"x":17.00000000000003,"y":795.2028985507245},"data":{"nodeID":"5HWIYkSdp12TH7pt","identifier":{"id":"SENS-Acanthocyte-over-RBC-count","source":"sour-candy"},"label":"Acanthocyte-over-RBC-count","short":"Acanthocyte","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":117.00000000000003,"y":795.2028985507245},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":127.27047464121942,"y":19.993945135586504,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.971014492753623,"y":19.993945135586504,"width":16,"height":16}]},"width":139,"height":56,"selected":false,"dragging":false},{"id":"nLFeCQNifqf3WhUe","type":"FindingComputeNode","position":{"x":34,"y":794.1159420289854},"data":{"nodeID":"nLFeCQNifqf3WhUe","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Acanthocytes Present","functionality":"reported","params":[{"value":0}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":609,"y":794.1159420289854},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":256.93215632784194,"y":27.964959628340125,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.971014492753623,"y":27.964959628340125,"width":16,"height":16}]},"width":269,"height":72,"selected":false,"dragging":false},{"id":"7gjPpASli4ccL7Mq","type":"WeightNode","position":{"x":10,"y":839.3236559533709},"data":{"nodeID":"7gjPpASli4ccL7Mq","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":0.5},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1060,"y":839.3236559533709},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":84.00900002834547,"y":39.98249035680065,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.011754831048162,"y":39.98249035680065,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"zIndex":1},{"id":"P9MuB4WgQCwFGNAr","type":"SensorNode","position":{"x":37,"y":642.4438433348037},"data":{"nodeID":"P9MuB4WgQCwFGNAr","identifier":{"id":"SENS-Serum-Ferritin-level","source":"sour-candy"},"label":"Serum Ferritin Level","short":"Serum Ferritin","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":137,"y":642.4438433348037},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":138.98488934046722,"y":19.991384892316884,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.011754831048162,"y":19.991384892316884,"width":16,"height":16}]},"width":151,"height":56,"selected":false,"dragging":false},{"id":"Pn68DuhmSkAIyZYi","type":"FindingComputeNode","position":{"x":37,"y":623.3071031549896},"data":{"nodeID":"Pn68DuhmSkAIyZYi","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Low Ferritin","functionality":"less_than","params":[{"value":"12","units":{"id":"12017","name":"micrograms per liter","short_1":"µg/l"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":612,"y":623.3071031549896},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":187.9711033358712,"y":28.00300000944849,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.011754831048162,"y":28.00300000944849,"width":16,"height":16}]},"width":200,"height":72,"selected":false,"dragging":false},{"id":"ilNqjse5fmKg6zT8","type":"WeightNode","position":{"x":33,"y":677.6142591271803},"data":{"nodeID":"ilNqjse5fmKg6zT8","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":2},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1083,"y":677.6142591271803},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":83.98189898176838,"y":39.98201375323236,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.035457255618493,"y":39.98201375323236,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"ghlqFx2MKGE9hYqI","type":"SensorNode","position":{"x":59,"y":507.0407779041069},"data":{"nodeID":"ghlqFx2MKGE9hYqI","identifier":{"id":"SENS-Serum-Iron-level","source":"sour-candy"},"label":"Serum Iron Level","short":"Serum Iron","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":159,"y":507.0407779041069},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":115.42141829796971,"y":20.010505559603306,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.996388465897031,"y":20.010505559603306,"width":16,"height":16}]},"width":127,"height":56,"selected":false,"dragging":false},{"id":"Lfb4fEI1o6kymnT8","type":"FindingComputeNode","position":{"x":54,"y":516.4449113884369},"data":{"nodeID":"Lfb4fEI1o6kymnT8","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Low Serum Iron","functionality":"noop","params":[{"value":0,"units":{"id":"1","name":"Unitless","short_1":"Unitless"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":629,"y":516.4449113884369},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":212.97761470740343,"y":28.006965449247545,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.996388465897031,"y":28.006965449247545,"width":16,"height":16}]},"width":225,"height":72,"selected":false,"dragging":false},{"id":"iFfFMg2TC7CdyHi3","type":"WeightNode","position":{"x":44,"y":547.59359170553},"data":{"nodeID":"iFfFMg2TC7CdyHi3","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":2},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1094,"y":547.59359170553},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":83.98189898176838,"y":39.98201375323236,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.035457255618493,"y":39.98201375323236,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"S6Xa2PyI3WjUweX1","type":"SensorNode","position":{"x":33,"y":389.6796971340237},"data":{"nodeID":"S6Xa2PyI3WjUweX1","identifier":{"id":"SENS-Total-Iron-Binding-Capacity","source":"sour-candy"},"label":"Total iron-binding capacity","short":"TIBC","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":133,"y":389.6796971340237},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":71.18754887358816,"y":20.010505559603306,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.996388465897031,"y":20.010505559603306,"width":16,"height":16}]},"width":83,"height":56,"selected":false,"dragging":false},{"id":"oFKC6mqdgnn2HWaP","type":"FindingComputeNode","position":{"x":52,"y":406.4242439667867},"data":{"nodeID":"oFKC6mqdgnn2HWaP","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"High TIBC","functionality":"greater_than_or_equal","params":[{"value":"450","units":{"id":"12018","name":"micrograms per deciliter","short_1":"µg/dl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":627,"y":406.4242439667867},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":188.0133008891143,"y":28.006965449247545,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-7.996388465897031,"y":28.006965449247545,"width":16,"height":16}]},"width":214,"height":72,"selected":false,"dragging":false},{"id":"sATJaGcIqTksHVbo","type":"WeightNode","position":{"x":40,"y":423.14755075425074},"data":{"nodeID":"sATJaGcIqTksHVbo","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":0.5},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1090,"y":423.14755075425074},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":83.98189898176838,"y":39.98201375323236,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.035457255618493,"y":39.98201375323236,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"kYtYC3kZph62mdTu","type":"SensorNode","position":{"x":10,"y":315.4911130389041},"data":{"nodeID":"kYtYC3kZph62mdTu","identifier":{"id":"SENS-Blood-Oxygen-Saturation","source":"sour-candy"},"label":"Blood Oxygen Saturation","short":"O2 Sat","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":110,"y":315.4911130389041},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":85.84020777169081,"y":19.98747639448746,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.00639408147741,"y":19.98747639448746,"width":16,"height":16}]},"width":98,"height":56},{"id":"Y9IbFCTQr14HOtKk","type":"FindingComputeNode","position":{"x":25,"y":308.5007303400084},"data":{"nodeID":"Y9IbFCTQr14HOtKk","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"Low Blood Oxygen Sat","functionality":"reported","params":[{"value":0,"units":{"id":"1","name":"Unitless","short_1":"Unitless"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":600,"y":308.5007303400084},"z":1000,"handleBounds":{"source":[{"id":null,"position":"right","x":258.750036260835,"y":27.99381811818855,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.00644643925373,"y":27.99381811818855,"width":16,"height":16}]},"width":271,"height":72,"selected":true,"dragging":false},{"id":"wFFLBZlxNfTsSbde","type":"WeightNode","position":{"x":41,"y":273.9320495688953},"data":{"nodeID":"wFFLBZlxNfTsSbde","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":1},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1091,"y":273.9320495688953},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":84.01014641598883,"y":40.00343541929282,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006341723701091,"y":40.00343541929282,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"NKo2orXMVglDRw80","type":"SensorNode","position":{"x":10,"y":245.14906884672197},"data":{"nodeID":"NKo2orXMVglDRw80","identifier":{"id":"SENS-Platelet-count","source":"sour-candy"},"label":"Platelet-count","short":"PLT","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":110,"y":245.14906884672197},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":62.99331613905635,"y":19.98747639448746,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006420260365571,"y":19.98747639448746,"width":16,"height":16}]},"width":75,"height":56},{"id":"nZW9SvlmQQ0veW8r","type":"FindingComputeNode","position":{"x":15,"y":190.15868614782624},"data":{"nodeID":"nZW9SvlmQQ0veW8r","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"High Patelet count","functionality":"greater_than_or_equal","params":[{"value":"450000","units":{"id":"125","name":"count per microliter","short_1":"/µl"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":590,"y":190.15868614782624},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":233.0724214514767,"y":27.99381811818855,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006341723701091,"y":27.99381811818855,"width":16,"height":16}]},"width":245,"height":72,"selected":false,"dragging":false},{"id":"lGYnLThyOcQGehgY","type":"WeightNode","position":{"x":39,"y":135.24796118453102},"data":{"nodeID":"lGYnLThyOcQGehgY","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":0.5},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1089,"y":135.24796118453102},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":84.01014641598883,"y":40.00343541929282,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006341723701091,"y":40.00343541929282,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1},{"id":"EKxUsy9S3iCplLKg","type":"SensorNode","position":{"x":29,"y":152.1105996423049},"data":{"nodeID":"EKxUsy9S3iCplLKg","identifier":{"id":"SENS-Red-Cell-Distribution-Width","source":"sour-candy"},"label":"Red Cell Distribution Width","short":"RDW","weight":0},"parentNode":"sensorContainer","extent":"parent","positionAbsolute":{"x":129,"y":152.1105996423049},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":72.17210554788417,"y":19.98747639448746,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.00639408147741,"y":19.98747639448746,"width":16,"height":16}]},"width":84,"height":56,"selected":false,"dragging":false},{"id":"N6rKvq9OEJIRgLRM","type":"FindingComputeNode","position":{"x":24,"y":92.05289583567922},"data":{"nodeID":"N6rKvq9OEJIRgLRM","identifier":{"id":"BLANK_FINDING","source":"sour-candy"},"label":"COMPUTE NODE","short":"New Finding","weight":0,"operation":{"title":"High RDW","functionality":"greater_than_or_equal","params":[{"value":"15","units":{"id":"20","name":"Percentage","short_1":"%"}}]}},"parentNode":"findingContainer","extent":"parent","positionAbsolute":{"x":599,"y":92.05289583567922},"z":0,"handleBounds":{"source":[{"id":null,"position":"right","x":188.00776922646511,"y":27.99381811818855,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006341723701091,"y":27.99381811818855,"width":16,"height":16}]},"width":200,"height":72,"selected":false,"dragging":false},{"id":"6FeYxrzYEsCXzm0k","type":"WeightNode","position":{"x":49,"y":29.45536316125333},"data":{"nodeID":"6FeYxrzYEsCXzm0k","identifier":{"id":"WEIGHT_NODE_v01","source":"sour-candy"},"weight":0.5},"parentNode":"logicContainer","extent":"parent","positionAbsolute":{"x":1099,"y":29.45536316125333},"z":1,"handleBounds":{"source":[{"id":null,"position":"right","x":84.01014641598883,"y":40.00343541929282,"width":16,"height":16}],"target":[{"id":null,"position":"left","x":-8.006341723701091,"y":40.00343541929282,"width":16,"height":16}]},"width":96,"height":96,"selected":false,"dragging":false,"zIndex":1}],"edges":[{"source":"hiJqpvH5hk48vIEZ","sourceHandle":null,"target":"TLbDaPOs9CO4DbA7","targetHandle":null,"id":"reactflow__edge-hiJqpvH5hk48vIEZ-TLbDaPOs9CO4DbA7","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"EUnVZT4c9Ub19d0B","sourceHandle":null,"target":"e5q43Fpckvkz36UN","targetHandle":null,"id":"reactflow__edge-EUnVZT4c9Ub19d0B-e5q43Fpckvkz36UN","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"EUnVZT4c9Ub19d0B","sourceHandle":null,"target":"33DQTv4BkYYmwBQu","targetHandle":null,"id":"reactflow__edge-EUnVZT4c9Ub19d0B-33DQTv4BkYYmwBQu","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"zVHcsLQ1twf5cv0K","sourceHandle":null,"target":"EfLwQJuY3Heo1bpg","targetHandle":null,"id":"reactflow__edge-zVHcsLQ1twf5cv0K-EfLwQJuY3Heo1bpg","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"zVHcsLQ1twf5cv0K","sourceHandle":null,"target":"NJmkKTJBqxaInVup","targetHandle":null,"id":"reactflow__edge-zVHcsLQ1twf5cv0K-NJmkKTJBqxaInVup","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"EfLwQJuY3Heo1bpg","sourceHandle":null,"target":"2OP30MUc57jrcwKX","targetHandle":null,"id":"reactflow__edge-EfLwQJuY3Heo1bpg-2OP30MUc57jrcwKX","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"33DQTv4BkYYmwBQu","sourceHandle":null,"target":"2OP30MUc57jrcwKX","targetHandle":null,"id":"reactflow__edge-33DQTv4BkYYmwBQu-2OP30MUc57jrcwKX","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"NJmkKTJBqxaInVup","sourceHandle":null,"target":"fzMlRxUGtB3odOX2","targetHandle":null,"id":"reactflow__edge-NJmkKTJBqxaInVup-fzMlRxUGtB3odOX2","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"e5q43Fpckvkz36UN","sourceHandle":null,"target":"fzMlRxUGtB3odOX2","targetHandle":null,"id":"reactflow__edge-e5q43Fpckvkz36UN-fzMlRxUGtB3odOX2","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"2OP30MUc57jrcwKX","sourceHandle":null,"target":"9DYkQJccdNqQSSLy","targetHandle":null,"id":"reactflow__edge-2OP30MUc57jrcwKX-9DYkQJccdNqQSSLy","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"9DYkQJccdNqQSSLy","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-9DYkQJccdNqQSSLy-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"5HWIYkSdp12TH7pt","sourceHandle":null,"target":"nLFeCQNifqf3WhUe","targetHandle":null,"id":"reactflow__edge-5HWIYkSdp12TH7pt-nLFeCQNifqf3WhUe","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"nLFeCQNifqf3WhUe","sourceHandle":null,"target":"7gjPpASli4ccL7Mq","targetHandle":null,"id":"reactflow__edge-nLFeCQNifqf3WhUe-7gjPpASli4ccL7Mq","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"7gjPpASli4ccL7Mq","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-7gjPpASli4ccL7Mq-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"P9MuB4WgQCwFGNAr","sourceHandle":null,"target":"Pn68DuhmSkAIyZYi","targetHandle":null,"id":"reactflow__edge-P9MuB4WgQCwFGNAr-Pn68DuhmSkAIyZYi","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"Pn68DuhmSkAIyZYi","sourceHandle":null,"target":"ilNqjse5fmKg6zT8","targetHandle":null,"id":"reactflow__edge-Pn68DuhmSkAIyZYi-ilNqjse5fmKg6zT8","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"ilNqjse5fmKg6zT8","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-ilNqjse5fmKg6zT8-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"ghlqFx2MKGE9hYqI","sourceHandle":null,"target":"Lfb4fEI1o6kymnT8","targetHandle":null,"id":"reactflow__edge-ghlqFx2MKGE9hYqI-Lfb4fEI1o6kymnT8","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"Lfb4fEI1o6kymnT8","sourceHandle":null,"target":"iFfFMg2TC7CdyHi3","targetHandle":null,"id":"reactflow__edge-Lfb4fEI1o6kymnT8-iFfFMg2TC7CdyHi3","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"iFfFMg2TC7CdyHi3","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-iFfFMg2TC7CdyHi3-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"oFKC6mqdgnn2HWaP","sourceHandle":null,"target":"sATJaGcIqTksHVbo","targetHandle":null,"id":"reactflow__edge-oFKC6mqdgnn2HWaP-sATJaGcIqTksHVbo","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"sATJaGcIqTksHVbo","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-sATJaGcIqTksHVbo-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"S6Xa2PyI3WjUweX1","sourceHandle":null,"target":"oFKC6mqdgnn2HWaP","targetHandle":null,"id":"reactflow__edge-S6Xa2PyI3WjUweX1-oFKC6mqdgnn2HWaP","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"kYtYC3kZph62mdTu","sourceHandle":null,"target":"Y9IbFCTQr14HOtKk","targetHandle":null,"id":"reactflow__edge-kYtYC3kZph62mdTu-Y9IbFCTQr14HOtKk","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"Y9IbFCTQr14HOtKk","sourceHandle":null,"target":"wFFLBZlxNfTsSbde","targetHandle":null,"id":"reactflow__edge-Y9IbFCTQr14HOtKk-wFFLBZlxNfTsSbde","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"wFFLBZlxNfTsSbde","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-wFFLBZlxNfTsSbde-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"NKo2orXMVglDRw80","sourceHandle":null,"target":"nZW9SvlmQQ0veW8r","targetHandle":null,"id":"reactflow__edge-NKo2orXMVglDRw80-nZW9SvlmQQ0veW8r","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"nZW9SvlmQQ0veW8r","sourceHandle":null,"target":"lGYnLThyOcQGehgY","targetHandle":null,"id":"reactflow__edge-nZW9SvlmQQ0veW8r-lGYnLThyOcQGehgY","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"lGYnLThyOcQGehgY","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-lGYnLThyOcQGehgY-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"N6rKvq9OEJIRgLRM","sourceHandle":null,"target":"6FeYxrzYEsCXzm0k","targetHandle":null,"id":"reactflow__edge-N6rKvq9OEJIRgLRM-6FeYxrzYEsCXzm0k","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"EKxUsy9S3iCplLKg","sourceHandle":null,"target":"N6rKvq9OEJIRgLRM","targetHandle":null,"id":"reactflow__edge-EKxUsy9S3iCplLKg-N6rKvq9OEJIRgLRM","style":{"strokeWidth":2,"stroke":"#a1a1aa"}},{"source":"6FeYxrzYEsCXzm0k","sourceHandle":null,"target":"1ePdHT3I06tb06zw","targetHandle":null,"id":"reactflow__edge-6FeYxrzYEsCXzm0k-1ePdHT3I06tb06zw","style":{"strokeWidth":2,"stroke":"#a1a1aa"}}],"viewport":{"x":-150.7856573276091,"y":-62.04884335818036,"zoom":0.755497766457708}}`
		f, err := flowchartJsonStringToWorkflow(s)
		Expect(err).To(BeNil())
		Expect(f).NotTo(BeNil())
		Expect(len(f.Nodes)).To(Equal(37))
		Expect(len(f.Edges)).To(Equal(32))
		Expect(f.Nodes[0].Type).To(Equal(ContainerNode))
		Expect(f.Nodes[1].Type).To(Equal(ContainerNode))
		Expect(f.Nodes[2].Type).To(Equal(ContainerNode))
		Expect(f.Nodes[3].Type).To(Equal(ContainerNode))
		Expect(f.Nodes[4].Type).To(Equal(SensorNode))
		Expect(f.Nodes[5].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[6].Type).To(Equal(SensorNode))
		Expect(f.Nodes[7].Type).To(Equal(SensorNode))
		Expect(f.Nodes[8].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[9].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[10].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[11].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[12].Type).To(Equal(LogicNode))
		Expect(f.Nodes[13].Type).To(Equal(LogicNode))
		Expect(f.Nodes[14].Type).To(Equal(DDxNode))
		Expect(f.Nodes[15].Type).To(Equal(WeightNode))
		Expect(f.Nodes[16].Type).To(Equal(SensorNode))
		Expect(f.Nodes[17].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[18].Type).To(Equal(WeightNode))
		Expect(f.Nodes[19].Type).To(Equal(SensorNode))
		Expect(f.Nodes[20].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[21].Type).To(Equal(WeightNode))
		Expect(f.Nodes[22].Type).To(Equal(SensorNode))
		Expect(f.Nodes[23].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[24].Type).To(Equal(WeightNode))
		Expect(f.Nodes[25].Type).To(Equal(SensorNode))
		Expect(f.Nodes[26].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[27].Type).To(Equal(WeightNode))
		Expect(f.Nodes[28].Type).To(Equal(SensorNode))
		Expect(f.Nodes[29].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[30].Type).To(Equal(WeightNode))
		Expect(f.Nodes[31].Type).To(Equal(SensorNode))
		Expect(f.Nodes[32].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[33].Type).To(Equal(WeightNode))
		Expect(f.Nodes[34].Type).To(Equal(SensorNode))
		Expect(f.Nodes[35].Type).To(Equal(FindingComputeNode))
		Expect(f.Nodes[36].Type).To(Equal(WeightNode))

		Expect(f.Edges[0].SourceID).To(Equal(f.Nodes[4].ID))
		Expect(f.Edges[0].TargetID).To(Equal(f.Nodes[5].ID))
	}
}

func Test_flowchartJsonStringToWorkflow_example_01(t *testing.T) {
	RegisterTestingT(t)

	{
		s := `{
			"nodes": [
					{
							"id": "start",
							"data": {
									"label": "Start"
							},
							"position": {
									"x": 0,
									"y": 50
							},
							"className": "rounded-full bg-zinc-200 font-medium",
							"style": {
									"width": 55,
									"height": 55
							},
							"sourcePosition": "right",
							"targetPosition": "right",
							"positionAbsolute": {
									"x": 0,
									"y": 50
							},
							"z": 0,
							"handleBounds": {
									"source": [
											{
													"id": null,
													"position": "right",
													"x": 52,
													"y": 24.5,
													"width": 6,
													"height": 6
											}
									],
									"target": [
											{
													"id": null,
													"position": "right",
													"x": 52,
													"y": 24.5,
													"width": 6,
													"height": 6
											}
									]
							},
							"width": 55,
							"height": 55
					},
					{
							"id": "c10i",
							"type": "FindingNode",
							"position": {
									"x": 109.375,
									"y": -32.75
							},
							"data": {
									"label": "Spherocytosis"
							},
							"positionAbsolute": {
									"x": 109.375,
									"y": -32.75
							},
							"z": 0,
							"handleBounds": {
									"source": [
											{
													"id": "a",
													"position": "right",
													"x": 149.433349609375,
													"y": 7,
													"width": 6,
													"height": 6
											},
											{
													"id": "b",
													"position": "right",
													"x": 149.433349609375,
													"y": 37,
													"width": 6,
													"height": 6
											}
									],
									"target": [
											{
													"id": null,
													"position": "left",
													"x": -4,
													"y": 25,
													"width": 6,
													"height": 6
											}
									]
							},
							"width": 151,
							"height": 56,
							"selected": false,
							"dragging": false
					},
					{
							"id": "ROhA",
							"type": "FindingNode",
							"position": {
									"x": 112.61552,
									"y": 121.71440000000001
							},
							"data": {
									"label": "Platelet"
							},
							"positionAbsolute": {
									"x": 112.61552,
									"y": 121.71440000000001
							},
							"z": 1000,
							"handleBounds": {
									"source": [
											{
													"id": "a",
													"position": "right",
													"x": 100.48665093749997,
													"y": 7.008765468749998,
													"width": 6,
													"height": 6
											},
											{
													"id": "b",
													"position": "right",
													"x": 100.48665093749997,
													"y": 36.99302906249999,
													"width": 6,
													"height": 6
											}
									],
									"target": [
											{
													"id": null,
													"position": "left",
													"x": -4.002058124999999,
													"y": 25.007623593749994,
													"width": 6,
													"height": 6
											}
									]
							},
							"width": 102,
							"height": 56,
							"selected": true,
							"dragging": false
					}
			],
			"edges": [
					{
							"id": "start-c10i",
							"source": "start",
							"target": "c10i",
							"label": "",
							"style": {},
							"markerStart": {
									"type": "arrow",
									"color": "#f00"
							},
							"markerEnd": {
									"type": "arrow",
									"color": "#f00"
							}
					},
					{
							"id": "start-ROhA",
							"source": "start",
							"target": "ROhA",
							"label": "",
							"style": {},
							"markerStart": {
									"type": "arrow",
									"color": "#f00"
							},
							"markerEnd": {
									"type": "arrow",
									"color": "#f00"
							}
					}
			],
			"viewport": {
					"x": 63.89673353909461,
					"y": 102.20897633744852,
					"zoom": 0.803755144032922
			}
	}
`
		f, err := flowchartJsonStringToWorkflow(s)
		Expect(err).To(BeNil())
		Expect(f).NotTo(BeNil())
		Expect(len(f.Nodes)).To(Equal(3))
		Expect(len(f.Edges)).To(Equal(2))
	}
}
