### Go Bowl simple rabbitmq Clean Architecture package
<img src="https://i.pinimg.com/564x/99/1c/6a/991c6a65099c731a52e926a8b6a5ddba.jpg" height="600px" width="300%" />
<br /> <br />
#### This package is a wrapper around [RabbitMQ](https://www.rabbitmq.com/) and [Amqp](https://github.com/rabbitmq/amqp091)
#### The syntax is very simple and can be easily understood

#### 👨‍💻 The codes below can help you get started:
* Install the package 

##### Run the below command at the root of your project to install the package:
    go get  go get github.com/kode-magic/go-bowl 
    go get  go get github.com/kode-magic/go-bowl/rabbit

##### Next is to create a configuration file where all rabbitmq configuration struct will be:
    touch config.go // place it to your desired location

##### Simple copy and paste the below code into the config file you created initially   

    package dice
    import "github.com/kode-magic/go-bowl/rabbit"

    func PublishConfigs() (*rabbit.MQConfig, *rabbit.ExchangeConfig, *rabbit.PublishConfig) {
        return &rabbit.MQConfig{
                Host:     "localhost",
                Port:     "5672",
                User:     "guest",
                Password: "guest",
            },
            &rabbit.ExchangeConfig{
                Exchange:   "school_space",
                Kind:       "fanout",
                Durable:    true,
                AutoDelete: false,
                Internal:   false,
                NoWait:     false,
            },
            &rabbit.PublishConfig{
                Exchange:    "school_space",
                RoutingKey:  "",
                Mandatory:   false,
                Immediate:   false,
                ContentType: "application/json",
            }
    }

    func ConsumeConfigs(consumer string) (*rabbit.MQConfig, *rabbit.ExchangeConfig, *rabbit.QueueConfig, *rabbit.ConsumerConfig) {
        return &rabbit.MQConfig{
                Host:     "localhost",
                Port:     "5672",
                User:     "guest",
                Password: "guest",
            }, &rabbit.ExchangeConfig{
                Exchange:   "school_space",
                Kind:       "fanout",
                Durable:    true,
                AutoDelete: false,
                Internal:   false,
                NoWait:     false,
            }, &rabbit.QueueConfig{
                Exchange:   "school_space",
                RoutingKey: "",
                QueueName:  "",
                Durable:    false,
                AutoDelete: false,
                Exclusive:  true,
                NoWait:     false,
            }, &rabbit.ConsumerConfig{
                Consumer:  consumer,
                AtoAck:    true,
                Exclusive: false,
                NoLocal:   false,
                NoWait:    false,
                Args:      nil,
            }
        
    }



