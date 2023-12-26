import logging
from random import randint
from opentelemetry import metrics


def metodo1():
    try:
        # Método 1
        meter = metrics.get_meter("cics.metricas")
        cics_total_transactions = meter.create_up_down_counter(
            name="cics.total_transactions",
            description="Quantidade de transações de um CICS",
        )
        res = randint(15000, 34000)
        cics_total_transactions.add(1, {"roll.value": res, "label2": 2})
        logging.getLogger().info(cics_total_transactions)
    except Exception as e:
        logging.getLogger().error(f"Erro no método 1: {e}")
    return


def metodo2():
    import logging
    from random import randint
    from opentelemetry import metrics
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import (
        ConsoleMetricExporter,
        PeriodicExportingMetricReader,
    )
    try:
        # Método 2
        metric_reader = PeriodicExportingMetricReader(ConsoleMetricExporter())
        provider = MeterProvider(metric_readers=[metric_reader])
        # Sets the global default meter provider
        metrics.set_meter_provider(provider)
        # Creates a meter from the global meter provider
        meter_2 = metrics.get_meter("cics.metrics")
        work_counter = meter_2.create_counter(
            "work.counter", unit="1", description="Counts the amount of work done"
        )
        res2 = randint(1000, 10000)
        work_counter.add(res2, {"work.type": "Teste"})
    except Exception as e:
        logging.getLogger().error(f"Erro no método 2: {e}")
    return


def metodo3():
    import logging
    from random import randint
    from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
    from opentelemetry.sdk.resources import SERVICE_NAME, Resource
    from opentelemetry import metrics
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import (
        PeriodicExportingMetricReader,
    )
    try:
        # Método 3 , com exporter
        resource = Resource(attributes={
            SERVICE_NAME: "your-service-name"
        })

        reader = PeriodicExportingMetricReader(
            OTLPMetricExporter(endpoint="localhost:4317")
        )
        provider = MeterProvider(resource=resource, metric_readers=[reader])
        metrics.set_meter_provider(provider)
        meter_3 = metrics.get_meter("cics.metricss")
        work_counter = meter_3.create_counter(
            "work.counter", unit="1", description="Counts the amount"
        )
        res3 = randint(1, 100)
        work_counter.add(res3, {"work.name": "Testeeeee"})
    except Exception as e:
        logging.getLogger().error(f"Erro no método 3: {e}")
    return


def metodo4():
    import logging
    from random import randint
    from opentelemetry.exporter.otlp.proto.http.metric_exporter import OTLPMetricExporter
    from opentelemetry.sdk.resources import SERVICE_NAME, Resource
    from opentelemetry import metrics
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import (
        PeriodicExportingMetricReader,
    )

    try:
        # Método 4 , com exporter http
        # Service name is required for most backends
        resource = Resource(attributes={
            SERVICE_NAME: "your-service-name"
        })

        reader = PeriodicExportingMetricReader(
            OTLPMetricExporter(endpoint="localhost:5555")
        )
        provider = MeterProvider(resource=resource, metric_readers=[reader])
        metrics.set_meter_provider(provider)
        meter_3 = metrics.get_meter("cics.metricss")
        work_counter = meter_3.create_counter(
            "work.counter", unit="1", description="Counts the amount"
        )
        res3 = randint(1, 100)
        work_counter.add(res3, {"work.name": "Testeeeee"})
    except Exception as e:
        logging.getLogger().error(f"Erro no método 4: {e}")
    return


def metodo5():
    import logging
    from opentelemetry.exporter.otlp.proto.http.metric_exporter import OTLPMetricExporter
    from opentelemetry import metrics
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import (
        PeriodicExportingMetricReader,
    )
    try:
        exporter = OTLPMetricExporter(endpoint="http://localhost:4317")
        reader = PeriodicExportingMetricReader(exporter, export_interval_millis=10000)
        metrics.set_meter_provider(MeterProvider(metric_readers=[reader]))

        meter = metrics.get_meter("cics.metrics4")

        counter = meter.create_up_down_counter(
            name="teste_metria_cics4",
            description="teste_metrica_cics44",
        )

        counter.add(1, {'label': "teste4"})
    except Exception as e:
        logging.getLogger().error(f"Erro no método 4: {e}")

    return


def metodo6():
    from opentelemetry.metrics import get_meter_provider, set_meter_provider
    from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader

    exporter = OTLPMetricExporter(endpoint="localhost:4317",insecure=True)
    # exporter = OTLPMetricExporter(insecure=True)
    reader = PeriodicExportingMetricReader(exporter)
    provider = MeterProvider(metric_readers=[reader])
    set_meter_provider(provider)

    meter = get_meter_provider().get_meter("nome-metrica")
    counter = meter.create_counter("primeiro_contador")
    counter.add(1)
    return
metodo6()

def metodo7():
    from opentelemetry import metrics
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import (
        ConsoleMetricExporter,
        PeriodicExportingMetricReader,
    )
    metric_reader = PeriodicExportingMetricReader(ConsoleMetricExporter())
    provider = MeterProvider(metric_readers=[metric_reader])
    metrics.set_meter_provider(provider)
    meter = metrics.get_meter("my.meter.name")
    work_counter = meter.create_counter(
        "work.counter", unit="1", description="Counts the amount of work done"
    )
    work_counter.add(1, {"work.type": "test-work"})
    return


def envia_evento():
    from opentelemetry import trace
    current_span = trace.get_current_span()
    current_span.add_event("Alerta no CICS")


def metodo8():
    # Com HTTP ao invés de GRPC
    from opentelemetry.metrics import get_meter_provider, set_meter_provider
    from opentelemetry.exporter.otlp.proto.http.metric_exporter import OTLPMetricExporter
    from opentelemetry.sdk.metrics import MeterProvider
    from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader

    exporter = OTLPMetricExporter(endpoint="my_otlp_endpoint:4317")
    reader = PeriodicExportingMetricReader(exporter)
    provider = MeterProvider(metric_readers=[reader])
    set_meter_provider(provider)
    meter = get_meter_provider().get_meter("nome-metrica")
    counter = meter.create_counter("primeiro_contador")
    counter.add(1)
    return