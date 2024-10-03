package jp.ac.tohoku.qse.takahashi.AtcSimulator.domain.model.entity.fix.util;

public class RoutePoint {
    private String name;
    private double latitude;
    private double longitude;
    private String type; // "waypoint" or "radio_navigation_aid"

    public RoutePoint(String name, double latitude, double longitude, String type) {
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
        this.type = type;
    }
}
