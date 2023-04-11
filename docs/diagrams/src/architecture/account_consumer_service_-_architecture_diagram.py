import os, sys
from urllib.request import urlretrieve

os.chdir(os.path.dirname(sys.argv[0]))

from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.compute import Pod
from diagrams.onprem.queue import Kafka
from diagrams.custom import Custom
from diagrams.programming.language import Go

with Diagram("account consumer service", show = False, direction="TB"):
    blueline=Edge(color="blue",style="bold")
    darkOrange=Edge(color="darkOrange",style="bold")
    blackline=Edge(color="black",style="bold")
    scylladb_url = "https://upload.wikimedia.org/wikipedia/en/2/20/Scylla_the_sea_monster.png"
    scylladb_icon = "rabbitmq.png"
    urlretrieve(scylladb_url, scylladb_icon)

    with Cluster("account-consumer-pod"):
        consumerPod=Pod("account-consumer-pod")

    with Cluster("external"):
       consumerCreateKafka=Kafka("account-createorupdate")
       consumerUpdateKafka=Kafka("account-createorupdate-dlq") 
       consumerCreateKafkaDlq=Kafka("account-delete")
       consumerUpdateKafkaDlq=Kafka("account-delete-dlq") 

    with Cluster("internal"):
       accountAvros=Go("account-toolkit")

    with Cluster("scyllaDb"):
       accountDatabase=Custom("account-database",scylladb_icon)

    consumerPod - darkOrange >> consumerCreateKafka
    consumerPod - darkOrange >> consumerUpdateKafka
    consumerPod - darkOrange >> consumerCreateKafkaDlq
    consumerPod - darkOrange >> consumerUpdateKafkaDlq
    consumerPod - blackline >> accountAvros
    consumerPod - blueline >> accountDatabase