import os, sys

os.chdir(os.path.dirname(sys.argv[0]))

from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.compute import Pod
from diagrams.onprem.queue import Kafka
from diagrams.aws.database import Aurora
from diagrams.k8s.network import Service


with Diagram("account consumer service", show = False):
    blueline=Edge(color="blue",style="bold")
    darkOrange=Edge(color="darkOrange",style="bold")
    blackline=Edge(color="black",style="bold")

    with Cluster("account-consumer-pod"):
        consumerPod=Pod("account-consumer-pod")

    with Cluster("external"):
       consumerCreateKafka=Kafka("account-create")
       consumerUpdateKafka=Kafka("account-update") 
       accountAvros=Service("account-toolkit")  

    with Cluster("scyllaDb"):
       accountDatabase=Aurora("account-database")

    consumerPod - darkOrange >> consumerCreateKafka
    consumerPod - darkOrange >> consumerUpdateKafka
    consumerPod - blackline >> accountAvros
    consumerPod - blueline >> accountDatabase