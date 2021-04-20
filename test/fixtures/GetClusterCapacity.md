GetClusterCapacity
You can use the GetClusterCapacity to return high-level capacity measurements for an entire
storage cluster. This method returns fields that you can use to calculate the efficiency rates shown in
the Element web UI. You can use the efficiency calculations in scripts to return the efficiency rates
for thin provisioning, deduplication, compression, and overall efficiency.
Efficiency calculations
Use the following equations to calculate thin provisioning, deduplication, and compression. These
equations apply to Element 8.2 and later.
• thinProvisioningFactor = (nonZeroBlocks + zeroBlocks) / nonZeroBlocks
Cluster API methods | 127
• deDuplicationFactor = (nonZeroBlocks + snapshotNonZeroBlocks) / uniqueBlocks
• compressionFactor = (uniqueBlocks * 4096) / (uniqueBlocksUsedSpace * 0.93)
Overall efficiency rate calculation
Use the following equation to calculate overall cluster efficiency using the results of the thin
provisioning, deduplication, and compression efficiency calculations.
• efficiencyFactor = thinProvisioningFactor * deDuplicationFactor * compressionFactor