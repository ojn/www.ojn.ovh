title: ZFS - OpenZFS file system
date: 2026-07-25
author: ojn, wikipedia
description: This post is about ZFS

# ZFS - OpenZFS

OpenZFS is an open-source implementation of the ZFS file system and volume manager initially developed by Sun Microsystems for the Solaris operating system, and is now maintained by the OpenZFS Project. Similar to the original ZFS, the implementation supports features like data compression, data deduplication, copy-on-write clones, snapshots, RAID-Z, and virtual devices that can create filesystems that span multiple disks.

One of the main capabilities of OpenZFS is self-healing. The file system can detect and correct errors while in use, without the need for a dedicated file system checker. This feature makes it suitable for mission-critical applications that require high availability.

OpenZFS is mainly used in enterprise and data center environments, as well as consumer devices like network-attached storage (NAS) devices, where data reliability and safety is essential. While initially designed for Solaris, development has since focused on Linux, while ports exist for various BSD distributions and macOS. Unlike Oracle ZFS, OpenZFS is licensed under the Common Development and Distribution License (CDDL), enabling both open-source and commercial use of the file system.

Founding members of OpenZFS include Matt Ahrens, one of the main architects of ZFS. In 2020, the codebases of OpenZFS and ZFS on Linux, a kernel module allowing ZFS to be used on Linux, were merged and released as OpenZFS 2.0, allowing other non-Linux operating systems to receive the various improvements that the Linux driver had incorporated over time.
