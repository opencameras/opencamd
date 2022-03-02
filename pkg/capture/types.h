typedef struct format {
	int codec;
	unsigned int width;
	unsigned int height;
	float max_frame_rate;
	float min_frame_rate;
} capture_format;

typedef struct device {
    const char *name;
	const char *model;
	const char *manufacturer;
	int format_count;
	capture_format **formats;
} capture_device;

typedef struct devices {
	int count;
	capture_device **devices;
} capture_devices;