
schema "_schema_" {
  table "sensor_readings" {
    column "upper4" {
      type = "integer[4]"
    }
    column "lower4" {
      type = "integer[4]"
    }
    column "sent5m" {
      type = "double precision[]"
    }
  }
}
