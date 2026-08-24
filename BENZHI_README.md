基于 Go 实现的泥炭地温室气体通量监测 Web 项目，一款野外作业后端服务，完成采样活动、闭合式气室观测、读数采集、校准、通量计算、质量评估与发布管理。

# Mireflux Container

Build the field operations service with `build_benzhi_docker.sh mireflux-run linux/amd64`.
The service persists its SQLite database at the `MIREFLUX_DB` path supplied at runtime.
